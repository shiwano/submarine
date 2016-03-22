module Build
  module Environment
    class << self
      def env
        (ENV['SUBMARINE_ENV'] || 'development').downcase
      end

      def version
        ENV['SUBMARINE_VERSION'] || '1.0.0'
      end

      def unity_path
        ENV['SUBMARINE_UNITY_PATH'] || '/Applications/Unity/Unity.app/Contents/MacOS/Unity'
      end

      [:production, :development].each do |name|
        define_method("#{name}?") do
          env == name.to_s
        end
      end
    end
  end
end
