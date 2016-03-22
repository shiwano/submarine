require './tools/build/environment'
require './tools/build/client_builder'

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
end
