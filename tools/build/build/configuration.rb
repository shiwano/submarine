require 'yaml'

module Build
  module Configuration
    class << self
      def config
        @config ||= open("config.#{Environment.env}.yml") {|f| YAML.load(f) }
      end

      def client
        config['client']
      end

      def build
        config['build']
      end

      def build_ios
        config['build']['ios']
      end
    end
  end
end
