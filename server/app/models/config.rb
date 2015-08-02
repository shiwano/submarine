require 'json'

class Config
  open("#{Rails.root}/../client/Assets/Resources/Config/Config.#{Rails.env}.json") do |file|
    json_data = JSON.load(file)
    config = TyphenApi::Model::Submarine::Config.new(json_data)

    config.attributes.each_key do |key|
      define_singleton_method(key) do |*args|
        config.send(key, *args)
      end
    end
  end
end
