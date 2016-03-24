require 'rails_helper'

RSpec.describe 'JoinIntoRoom', type: :request do
  describe 'POST /join_into_room', with_login: true do
    let(:room) { create(:room) }
    let(:params) { { room_id: room.id } }

    context 'with a valid request' do
      before do
        post(join_into_room_path, params)
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return the joined room' do
        expect(response_json[:room][:id]).to eq current_user.reload.room.id
      end
    end

    context 'with a no-existing room_id' do
      let(:params) { { room_id: -1 } }

      it 'should not work' do
        expect { post(join_into_room_path, params) }.to raise_error GameError::RoomNotFound
      end
    end
  end
end
