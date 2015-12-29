require 'rails_helper'

RSpec.describe "Battle::FindRoom", type: :request do
  describe "POST /battle/find_room" do
    let(:room) { create(:room) }
    let(:params) { { room_id: room.id } }

    before do
      post(battle_find_room_path, params)
    end

    context 'with a valid request' do
      it 'should work' do
        expect(response).to have_http_status(200)
      end
      it 'should return the requested room' do
        expect(response_json[:room][:id]).to eq params[:room_id]
      end
    end

    context 'with a no-existing room_id' do
      let(:params) { { room_id: -1 } }

      it 'should return nil as the requested room' do
        expect(response_json[:room]).to eq nil
      end
    end
  end
end
