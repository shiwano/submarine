require File.expand_path('../boot', __FILE__)

require 'rails/all'

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module SubmarineApi
  class Application < Rails::Application
    # Settings in config/environments/* take precedence over those specified here.
    # Application configuration should go into files in config/initializers
    # -- all .rb files in that directory are automatically loaded.

    # Set Time.zone default to the specified zone and make Active Record auto-convert to this zone.
    # Run "rake -D time" for a list of tasks for finding time zone names. Default is UTC.
    # config.time_zone = 'Central Time (US & Canada)'

    # The default locale is :en and all translations from config/locales/*.rb,yml are auto loaded.
    # config.i18n.load_path += Dir[Rails.root.join('my', 'locales', '*.{rb,yml}').to_s]
    # config.i18n.default_locale = :de

    # Do not swallow errors in after_commit/after_rollback callbacks.
    config.active_record.raise_in_transactional_callbacks = true

    config.paths["config/routes.rb"] << "#{Rails.root}/lib/typhen_api/typhen_api/routes.rb"
    config.autoload_paths << "#{Rails.root}/lib/typhen_api"
    config.assets.enabled = false

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
