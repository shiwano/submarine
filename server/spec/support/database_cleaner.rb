require 'database_cleaner'

RSpec.configure do |config|
  config.before(:suite) do
    DatabaseCleaner[:redis].db = "redis://#{Settings.redis[:host]}:#{Settings.redis[:port]}/#{Settings.redis[:db]}"
    DatabaseCleaner[:active_record].strategy = :transaction
    DatabaseCleaner.clean_with :truncation
  end

  config.before(:each) do
    DatabaseCleaner.start
  end

  config.after(:each) do
    DatabaseCleaner.clean
  end
end
