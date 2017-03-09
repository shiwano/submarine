module TyphenApiRespondable
  RENDER_FORMAT = :msgpack

  def params
    @typhen_api_params ||= self.class::RequestType.new(super.except(:controller, :action).to_unsafe_h)
  end

  def render(response_body, error: false)
    if error
      super RENDER_FORMAT => to_serializable_object(self.class::ErrorType.new(response_body)), status: 500
    elsif self.class::ResponseType.present?
      super RENDER_FORMAT => to_serializable_object(self.class::ResponseType.new(response_body))
    else
      super RENDER_FORMAT => response_body
    end
  end

  private

  def to_serializable_object(target)
    return target unless target.respond_to?(:to_h)
    target_as_hash = target.to_h
    return nil if target_as_hash.keys.size == 0

    target_as_hash.each_with_object({}) do |(k, v), h|
      if v.is_a? Array
        h[k] = v.map {|v| to_serializable_object(v) }
      else
        h[k] = to_serializable_object(v)
      end
    end
  end
end
