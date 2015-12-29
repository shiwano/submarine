require 'rails_helper'

RSpec.describe Battle::CloseRoomController, type: :controller do
  context 'POST service' do
    let(:room) { create(:room) }
    let(:params) { { room_id: room.id } }

    describe 'target_room' do
      subject { assigns(:target_room) }

      it 'should return a Room' do
        post :service, params
        expect(subject).to be_a_kind_of Room
      end
    end
  end
end
