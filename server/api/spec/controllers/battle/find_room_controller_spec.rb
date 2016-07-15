require 'rails_helper'

RSpec.describe Battle::FindRoomController, type: :controller do
  context 'POST service' do
    let(:room) { create(:room) }
    let(:params) { { room_id: room.id } }

    describe '#room' do
      subject { assigns(:room) }

      it 'should return a Room' do
        post :service, params: params
        expect(subject).to be_a_kind_of Room
      end
    end
  end
end
