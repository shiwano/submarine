module TyphenApiRespondable
  def params
    @typhen_api_params ||= self.class::RequestType.new(
      super.except(:controller, :action).to_unsafe_h
    )
  end

  def render_response(raw_response)
    if self.class::ResponseType.present?
      response = self.class::ResponseType.new(raw_response)
      render :json => response
    else
      render :json => raw_response
    end
  end

  def render_error(raw_response, status)
    response = self.class::ErrorType.new(raw_response)
    render :json => response, :status => status
  end
end
