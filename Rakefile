require './tools/build/build'
require 'fileutils'
include FileUtils

namespace :gen do
  desc 'Generate code from the contract'
  task :contract do
    sh 'typhen'
  end

  desc 'Generate light map JSON files'
  task :lightmap do
    cd 'tools/lightmap' do
      sh 'go build && ./lightmap 5 20'
    end
  end
end

namespace :build do
  desc 'Build the client for iOS'
  task :ios do
    Build::ClientBuilder.build(:ios)
  end

  desc 'Build the client for Android'
  task :android do
    Build::ClientBuilder.build(:android)
  end

  desc 'Generate the config file for the client'
  task :generate_config_for_client do
    Build::ClientBuilder.generate_config_for_client
  end
end
