class ApplicationController < ActionController::API
  before_action :authentication_with_access_token, unless: :no_authentication_required?
  after_action :generate_access_token_if_needed, unless: :no_authentication_required?
  rescue_from StandardError, with: :render_error unless Rails.env.test?

  attr_reader :current_user

  protected

  def set_access_token(access_token)
    response.headers['X-Set-Access-Token'] = access_token
  end

  private

  def authentication_with_access_token
    access_token = request.headers['X-Access-Token']
    raise GameError::NotAuthenticated.new('No access token') if access_token.nil?
    @current_user = User.find_by_access_token(access_token)
    raise GameError::NotAuthenticated.new('Invalid access token') if current_user.nil?
  end

  def generate_access_token_if_needed
    if current_user.access_token.expires_soon?
      set_access_token(current_user.generate_access_token!)
    end
  end

  def render_error(e)
    backtrace_text = simple_backtrace_text(e)
    logger.error([e.message, backtrace_text].flatten.join("\n"))
    response_body = { code: 500, name: e.class.name, message: e.message }
    render response_body, error: true
  end

  def simple_backtrace_text(e)
    bc = ActiveSupport::BacktraceCleaner.new
    bc.add_filter   {|line| line.gsub(Rails.root.to_s, '') }
    bc.add_silencer {|line| line =~ /ruby\/gems|webrick/ }
    bc.clean(e.backtrace).join("\n")
  end
end
