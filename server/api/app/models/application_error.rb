module ApplicationError
  class LoginFailed < StandardError; end
  class SessionServerDown < StandardError; end
  class NotAuthenticated < StandardError; end

  class RoomIsFull < StandardError; end
  class RoomAlreadyJoined < StandardError; end
  class RoomNotFound < StandardError; end
end
