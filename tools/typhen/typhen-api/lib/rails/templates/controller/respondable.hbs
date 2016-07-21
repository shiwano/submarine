module TyphenApiRespondable
  def params
    @typhen_api_params ||= self.class::RequestType.new(super.except(:controller, :action).to_unsafe_h)
  end

  def render(response_body, error: false)
    if error
      super json: self.class::ErrorType.new(response_body), status: 500
    elsif self.class::ResponseType.present?
      super json: self.class::ResponseType.new(response_body)
    else
      super json: response_body
    end
  end
end
