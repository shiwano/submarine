# Be sure to restart your server when you modify this file.

# Rails.application.config.session_store :cookie_store, key: '_submarine_api_session'

Rails.application.config.session_store :redis_session_store, {
  key: '_submarine_api_session',
  serializer: :json,
  on_redis_down: ->(e, env, sid) { raise ApplicationError::SessionServerDown.new(e) },
  redis: {
    db: 3,
    expire_after: 120.minutes,
    key_prefix: 'submarine:session:'
  }
}
