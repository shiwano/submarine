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
      @config = open("tools/build/config.#{Environment.env}.yml") {|f| YAML.load(f) }
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
      rm_rf 'client/iOSXCodeProject'
      build_with_unity

      cd 'client/iOSXCodeProject' do
        sh <<-EOS
          xcodebuild -version
          xcodebuild clean
          xcodebuild -sdk iphoneos CODE_SIGN_IDENTITY="iPhone Distribution: #{build_config['ios']['code_sign_identity']}"
        EOS

        app_path = "build/Release-iphoneos/#{build_config['product_name'].downcase}.app"
        unless File.exist?("#{app_path}/ResourceRules.plist")
          cp "#{@workspace}/tools/build/ResourceRules.plist", app_path
        end

        sh <<-EOS
          xcrun -sdk iphoneos "PackageApplication" "#{app_path}" \
            -o "#{@workspace}/client/build_#{Environment.version}.ipa" \
            --sign "#{build_config['ios']['code_sign_identity']}" \
            --embed "#{@workspace}/#{build_config['ios']['provisioning_profile']}"
        EOS
      end
    end

    def build_for_android
      build_with_unity
      mv 'client/build.apk', "client/build_#{Environment.version}.apk"
    end
  end
end
