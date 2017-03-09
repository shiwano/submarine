if Module::const_defined?(:MessagePack)
  Mime::Type.register "application/x-msgpack", :msgpack

  ActionDispatch::Request.parameter_parsers[:msgpack] = -> (raw_post) do
    MessagePack.unpack(raw_post)
  end

  ActionController::Renderers.add :msgpack do |body, options|
    self.content_type ||= Mime[:msgpack]
    body.respond_to?(:to_msgpack) ? body.to_msgpack(options) : body
  end
end
