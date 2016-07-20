require 'rails_helper'

RSpec.describe 'CreateRoom', type: :request do
  describe 'POST /create_room', with_login: true do
    context 'with a valid request' do
      before do
        post create_room_path
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return a reasponse that includes a user' do
        expect(parsed_response.room.id).to eq current_user.reload.room.id
      end
    end

    context 'with the user has already a room' do
      before do
        current_user.create_room!
      end

      it 'should not work' do
        expect { post create_room_path }.to raise_error GameError::RoomAlreadyJoined
      end
    end
  end
end
