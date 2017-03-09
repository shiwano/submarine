module ResponseTestHelpers
  module Integration
    def parsed_response
      @parsed_response ||= Hashie::Mash.new(response.parsed_body)
    end
  end
end

RSpec.configure do |config|
  config.include ResponseTestHelpers::Integration, type: :request
end
