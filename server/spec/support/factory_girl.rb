RSpec.configure do |config|
  config.include FactoryGirl::Syntax::Methods

  config.before :suite do
    begin
      DatabaseRewinder.start
      FactoryGirl.lint
    ensure
      DatabaseRewinder.clean
    end
  end
end
