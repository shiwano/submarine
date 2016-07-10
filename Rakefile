require './tools/build/build'

desc 'Generate code from the contract'
task :gen do
  sh 'typhen'
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
