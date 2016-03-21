require './tools/build/client_builder'

ENV_NAMES = [
  :development,
  :production,
]

desc 'Generate code from the contract'
task :gen do
  sh 'typhen'
end

ENV_NAMES.each do |env|
  namespace env do
    namespace :build do
      desc 'Build the client for iOS'
      task :ios do
        ClientBuilder.build(env, :ios)
      end

      desc 'Build the client for Android'
      task :android do
        ClientBuilder.build(env, :android)
      end
    end
  end
end
