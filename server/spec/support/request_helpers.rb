module Requests
  module JsonHelpers
    def response_json
      @response_json ||= JSON.parse(response.body)
    end
  end
end

RSpec.configure do |config|
  config.include Requests::JsonHelpers, type: :request
end
