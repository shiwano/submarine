class GameConfig
  open("#{Rails.root}/../../config.#{Rails.env.test? ? 'example' : Rails.env}.yml") do |file|
    hash = YAML.load(file)
    config = Hashie::Mash.new(hash).server

    config.keys.each do |key|
      define_singleton_method(key) do |*args|
        config.send(key, *args)
      end
    end
  end
end
