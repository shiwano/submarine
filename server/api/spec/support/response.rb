module ResponseTestHelpers
  module Integration
    def parsed_response
      @parsed_response ||= Hashie::Mash.new(JSON.parse(response.body, :symbolize_names => true))
    end
  end
end

RSpec.configure do |config|
  config.include ResponseTestHelpers::Integration, type: :request
end
