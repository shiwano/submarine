require 'yaml'
require 'json'
require 'fileutils'
include FileUtils

module Build
  class ClientBuilder
    class << self
      def build(target)
        self.new(target).build
      end
    end

    def initialize(target)
      raise "Unsupported target: #{target}" unless [:ios, :android].include?(target)
      @workspace = Dir.pwd
      @target = target
      @config = open("config.#{Environment.env}.yml") {|f| YAML.load(f) }
    end

    def build_config
      @config['build']
    end

    def build
      make_client_config
      case @target
      when :ios     then build_for_ios
      when :android then build_for_android
      end
    end

    def make_client_config
      open("client/Assets/Resources/Config/Config.json", 'w') do |file|
        @config['client_config']['version'] = Environment.version
        JSON.dump(@config['client_config'], file)
      end
    end

    def build_with_unity
      sh <<-EOS
        #{Environment.unity_path} \
          -quit \
          -batchmode \
          -executeMethod BuildScript.ExecuteViaCommandLine \
          -projectPath #{@workspace}/client \
          -logFile #{@workspace}/client/build.log \
          "buildTarget=#{@target}" \
          "bundleVersion=#{Environment.version}" \
          "bundleIdentifier=#{build_config['bundle_identifier']}" \
          "productName=#{build_config['product_name']}"
      EOS
    end

    def build_for_ios
      output_path = "#{@workspace}/client/build_#{Environment.version}.ipa"
      rm output_path
      rm_rf 'client/iOSXCodeProject'
      build_with_unity

      cd 'client/iOSXCodeProject' do
        sh <<-EOS
          xcodebuild -version
          xcodebuild clean

          xcodebuild -configuration Release -scheme Unity-iPhone \
            -archivePath "#{build_config['product_name'].downcase}.xcarchive" \
            PROVISIONING_PROFILE="#{build_config['ios']['provisioning_profile']}" \
            CODE_SIGN_IDENTITY="iPhone Distribution: #{build_config['ios']['code_sign_identity']}" \
            archive

          xcodebuild -exportArchive -exportFormat IPA \
            -archivePath "#{build_config['product_name'].downcase}.xcarchive" \
            -exportPath "#{output_path}" \
            PROVISIONING_PROFILE="#{build_config['ios']['provisioning_profile']}" \
            CODE_SIGN_IDENTITY="iPhone Distribution: #{build_config['ios']['code_sign_identity']}"
        EOS
      end
    end

    def build_for_android
      build_with_unity
      mv 'client/build.apk', "client/build_#{Environment.version}.apk"
    end
  end
end
