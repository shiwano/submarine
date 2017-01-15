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
      sh 'go build && ./lightmap 20 80'
    end
  end

  desc 'Generate all configuration files'
  task :configs do
    Build::ClientBuilder.generate_config
    Build::ServerBuilder.generate_all_configs
  end
end

namespace :build do
  namespace :client do
    desc 'Build the client for iOS'
    task :ios do
      Build::ClientBuilder.build(:ios)
    end

    desc 'Build the client for Android'
    task :android do
      Build::ClientBuilder.build(:android)
    end
  end

  namespace :server do
    Build::ServerBuilder::TARGETS.each do |target|
      desc "Bulid a docker image of the #{target} server"
      task target do
        Build::ServerBuilder.build(target)
      end
    end
  end
end
