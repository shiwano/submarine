module TyphenApiRespondable
  extend ActiveSupport::Concern

  def params_with_validation
    @typhen_api_params ||= self.class::RequestType.new(params_without_validation)
  end

  def render_response(response)
    render :json => response
  end

  def render_response_with_validation(raw_response)
    if self.class::ResponseType.present?
      response = self.class::ResponseType.new(raw_response)
      render_response_without_validation(response)
    else
      render_response_without_validation({})
    end
  end

  def render_error(response, status)
    render :json => response, :status => status
  end

  def render_error_with_validation(raw_response, status)
    response = self.class::ErrorType.new(raw_response)
    render_error_without_validation(response, status)
  end

  included do
    alias_method_chain(:params, :validation)
    alias_method_chain(:render_response, :validation) unless Rails.env.production?
    alias_method_chain(:render_error, :validation) unless Rails.env.production?
  end
end
