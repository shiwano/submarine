require 'rails_helper'

RSpec.describe 'JoinIntoRoom', type: :request do
  describe 'POST /join_into_room' do
    let(:user) { create :user, :with_stupid_password }
    let(:room) { create :room }
    let(:params) { { room_id: room.id } }

    before do
      login_user(user)
    end

    context 'with a valid request' do
      before do
        post join_into_room_path, params
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return the joined room' do
        expect(response_json[:room][:id]).to eq user.reload.room.id
      end
    end

    context 'with a invalid room_id' do
      let(:params) { { room_id: -1 } }

      it 'should not work' do
        expect { post join_into_room_path, params }.to raise_error ApplicationError::RoomNotFound
      end
    end
  end
end
