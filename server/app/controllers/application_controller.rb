class ApplicationController < ActionController::Base
  # Prevent CSRF attacks by raising an exception.
  # For APIs, you may want to use :null_session instead.
  # protect_from_forgery with: :exception

  before_action :require_login_for_api
  rescue_from StandardError, with: :render_500

  def require_login_for_api
    require_login unless no_authentication_required?
  end

  def not_authenticated
    raise ApplicationError::NotAuthenticated.new('Not authenticated')
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
