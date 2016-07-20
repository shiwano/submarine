require 'rails_helper'

RSpec.describe "Battle::CloseRoom", type: :request do
  describe "POST /battle/close_room" do
    let(:room) { create(:room) }
    let(:request_params) { { room_id: room.id } }

    context 'with a valid request' do
      before do
        post battle_close_room_path, params: request_params
      end

      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return a empty response' do
        expect(parsed_response).to be_empty
      end
      it 'should destroy the room' do
        expect { room.reload }.to raise_error ActiveRecord::RecordNotFound
      end
    end

    context 'with a no-existing room_id' do
      let(:request_params) { { room_id: -1 } }

      it 'should not work' do
        expect {
          post battle_close_room_path, params: request_params
        }.to raise_error GameError::RoomNotFound
      end
    end
  end
end
