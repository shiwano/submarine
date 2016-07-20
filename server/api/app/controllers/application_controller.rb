class ApplicationController < ActionController::API
  before_action :authentication_with_access_token, unless: :no_authentication_required?
  rescue_from StandardError, with: :render_500

  attr_reader :current_user

  private

  def authentication_with_access_token
    access_token = request.headers['X-Access-Token']
    raise GameError::NotAuthenticated.new('No access token') if access_token.nil?
    @current_user = User.find_by_access_token(access_token)
    raise GameError::NotAuthenticated.new('Invalid access token') if current_user.nil?
  end

  def render_500(e)
    if Rails.env.test?
      raise e
    else
      backtrace_text = simple_backtrace_text(e)
      logger.error([e.message, backtrace_text].flatten.join("\n"))
      render_error({ code: 500, name: e.message, message: backtrace_text }, 500)
    end
  end

  def simple_backtrace_text(e)
    bc = ActiveSupport::BacktraceCleaner.new
    bc.add_filter   { |line| line.gsub(Rails.root.to_s, '') }
    bc.add_silencer { |line| line =~ /ruby\/gems|webrick/ }
    bc.clean(e.backtrace).join("\n")
  end
end
