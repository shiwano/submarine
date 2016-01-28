require 'rails_helper'

RSpec.describe GetRoomsController, type: :controller do
  context 'POST service', with_login: true do
    describe '#joinable_rooms' do
      it 'should receive Room.joinable' do
        expect(Room).to receive(:joinable).and_call_original
        post :service
      end
    end
  end
end
