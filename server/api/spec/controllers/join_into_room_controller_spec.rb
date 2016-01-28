require 'rails_helper'

RSpec.describe JoinIntoRoomController, type: :controller do
  context 'POST service', with_login: true do
    let(:room) { create(:room) }
    let(:params) { { room_id: room.id } }

    describe '#room' do
      subject { assigns(:room) }

      it 'should return a Room' do
        post :service, params
        expect(subject).to be_a_kind_of Room
      end

      it 'should call Room.joinable' do
        expect(Room).to receive(:joinable).and_call_original
        post :service, params
      end
    end
  end
end
