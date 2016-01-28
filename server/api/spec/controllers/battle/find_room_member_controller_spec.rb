require 'rails_helper'

RSpec.describe Battle::FindRoomMemberController, type: :controller do
  context 'POST service' do
    let(:room_member) { create(:room_member) }
    let(:params) { { room_key: room_member.room_key } }

    describe '#room_member' do
      subject { assigns(:room_member) }

      it 'should return the room member' do
        post :service, params
        expect(subject).to eq room_member
      end
    end
  end
end
