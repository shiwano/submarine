require 'yaml'
require 'json'
require 'fileutils'
include FileUtils

class ClientBuilder
  class << self
    def build(env, target)
      self.new(env, target).build
    end
  end

  def initialize(env, target)
    raise "Unsupported target: #{target}" unless [:ios, :android].include?(target)
    @workspace = Dir.pwd
    @env = env
    @target = target
    @unity_path = ENV['UNITY_PATH'] || '/Applications/Unity/Unity.app/Contents/MacOS/Unity'
    @version = ENV['BUILD_VERSION'] || '1.0.0'
    @config = open("tools/build/config.#{@env}.yml") {|f| YAML.load(f) }
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
    open("client/Assets/Resources/Config/Config.#{@env}.json", 'w') do |file|
      @config['client_config']['version'] = @version
      JSON.dump(@config['client_config'], file)
    end
  end

  def build_with_unity
    sh <<-EOS
      #{@unity_path} \
        -quit \
        -batchmode \
        -executeMethod BuildScript.ExecuteViaCommandLine \
        -projectPath #{@workspace}/client \
        -logFile #{@workspace}/client/build.log \
        "buildTarget=#{@target}" \
        "bundleVersion=#{@version}" \
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
        cp '../../tools/build/ResourceRules.plist', app_path
      end

      sh <<-EOS
        xcrun -sdk iphoneos "PackageApplication" "#{app_path}" \
          -o "#{@workspace}/client/build_#{@version}.ipa" \
          --sign "#{build_config['ios']['code_sign_identity']}" \
          --embed "#{@workspace}/#{build_config['ios']['provisioning_profile']}"
      EOS
    end
  end

  def build_for_android
    build_with_unity
    mv 'client/build.apk', "client/build_#{@version}.apk"
  end
end
