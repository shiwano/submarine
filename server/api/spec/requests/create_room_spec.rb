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
        expect(response_json[:room][:id]).to eq current_user.reload.room.id
      end
    end

    context 'with the user has already a room' do
      let(:current_user) { create(:user, :with_stupid_password, :with_room) }

      it 'should not work' do
        expect { post create_room_path }.to raise_error ApplicationError::RoomAlreadyJoined
      end
    end
  end
end
