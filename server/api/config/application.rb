require_relative 'boot'

require 'rails/all'

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module SubmarineApi
  class Application < Rails::Application
    # Settings in config/environments/* take precedence over those specified here.
    # Application configuration should go into files in config/initializers
    # -- all .rb files in that directory are automatically loaded.

    config.paths["config/routes.rb"] << "#{Rails.root}/lib/typhen_api/typhen_api/routes.rb"
    config.autoload_paths << "#{Rails.root}/lib/typhen_api"

    config.api_only = true
    config.debug_exception_response_format = :api

    config.generators do |g|
      g.assets false
      g.helper false
      g.template_engine false
      g.fixture_replacement :factory_girl, dir: 'spec/factories'
      g.test_framework :rspec,
        fixture: true,
        view_specs: false,
        routing_specs: false,
        helper_specs: false,
        integration_tool: false
    end
  end
end
