require 'json'
require 'fileutils'
require 'plist'
include FileUtils

module Build
  class ClientBuilder
    class << self
      def build(target)
        self.new(target).build
      end

      def generate_config
        mkdir 'client/Assets/Resources/Config'
        open('client/Assets/Resources/Config/Config.json', 'w') do |file|
          Configuration.client['version'] = Environment.version
          JSON.dump(Configuration.client, file)
        end
      end
    end

    def initialize(target)
      raise "Unsupported target: #{target}" unless [:ios, :android].include?(target)
      @workspace = Dir.pwd
      @target = target
    end

    def build
      self.class.generate_config
      case @target
      when :ios     then build_for_ios
      when :android then build_for_android
      end
    end

    private

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
          "bundleIdentifier=#{Configuration.build['bundle_identifier']}" \
          "productName=#{Configuration.build['product_name']}"
      EOS
    end

    def build_for_ios
      rm_f "client/build_#{Environment.version}.ipa"
      rm_rf 'client/iOSXCodeProject'
      build_with_unity

      open("#{@workspace}/client/iOSXCodeProject/exportOptions.plist", 'w') do |file|
        file.write Configuration.build_ios['export_options'].to_plist
      end

      cd 'client/iOSXCodeProject' do
        sh <<-EOS
          export CODE_SIGN_IDENTITY="#{Configuration.build_ios['code_sign_identity']}"
          export PROVISIONING_PROFILE="#{Configuration.build_ios['provisioning_profile']}"
          xcodebuild -version
          xcodebuild clean
          xcodebuild \
            -configuration Release \
            -scheme Unity-iPhone \
            -archivePath "#{Configuration.build['product_name'].downcase}.xcarchive" \
            archive
          xcodebuild \
            -exportArchive \
            -archivePath "#{Configuration.build['product_name'].downcase}.xcarchive" \
            -exportPath "#{@workspace}/client" \
            -exportOptionsPlist "#{@workspace}/client/iOSXCodeProject/exportOptions.plist"
        EOS
      end

      mv 'client/Unity-iPhone.ipa', "client/build_#{Environment.version}.ipa"
    end

    def build_for_android
      build_with_unity
      mv 'client/build.apk', "client/build_#{Environment.version}.apk"
    end
  end
end
