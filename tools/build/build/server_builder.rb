require 'json'
require 'fileutils'
include FileUtils

module Build
  class ServerBuilder
    class << self
      def build(target)
        self.new(target).build
      end

      def generate_all_configs
        ServerBuilder::TARGETS.each do |target|
          self.new(target).generate_config
        end
      end
    end

    TARGETS = [:api, :battle]

    def initialize(target)
      raise "Unsupported target: #{target}" unless TARGETS.include?(target)
      @workspace = Dir.pwd
      @target = target
    end

    def build
      generate_config
      build_with_docker
    end

    def generate_config
      open("#{server_directory}/config.json", 'w') do |file|
        JSON.dump(Configuration.server, file)
      end
    end

    private

    def server_directory
      "server/#{@target}"
    end

    def build_with_docker
      cd server_directory do
        sh "docker build -t submarine-#{Environment.env}-#{@target} --build-arg env=#{Environment.env} ."
      end
    end
  end
end
