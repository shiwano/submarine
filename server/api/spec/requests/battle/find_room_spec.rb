require 'rails_helper'

RSpec.describe "Battle::FindRoom", type: :request do
  describe "POST /battle/find_room" do
    let(:room) { create(:room) }
    let(:request_params) { { room_id: room.id } }

    before do
      post battle_find_room_path, params: request_params
    end

    context 'with a valid request' do
      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return the requested room' do
        expect(parsed_response.room.id).to eq request_params[:room_id]
      end
    end

    context 'with a no-existing room_id' do
      let(:request_params) { { room_id: -1 } }

      it 'should return nil as the requested room' do
        expect(parsed_response.room).to eq nil
      end
    end
  end
end
