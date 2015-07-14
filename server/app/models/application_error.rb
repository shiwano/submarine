module ApplicationError
  class LoginFailed < StandardError; end
  class SessionServerDown < StandardError; end
  class NotAuthenticated < StandardError; end
end
