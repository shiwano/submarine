class GameConfig
  open("#{Rails.root}/config#{Rails.env.test? ? '.test' : ''}.json") do |file|
    hash = JSON.load(file)
    config = TyphenApi::Model::Submarine::Configuration::Server.new(hash)

    config.attributes.keys.each do |key|
      define_singleton_method(key) do |*args|
        config.send(key, *args)
      end
    end
  end
end
