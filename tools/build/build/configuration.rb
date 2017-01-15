require 'yaml'

module Build
  module Configuration
    class << self
      def path
        "config.#{Environment.env}.yml"
      end

      def config
        @config ||= open(path) {|f| YAML.load(f) }
      end

      def client
        config['client']
      end

      def server
        config['server']
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
