require 'rails_helper'

RSpec.describe GetRoomsController, type: :controller do
  context 'POST service' do
    before do
      user = create :user, :with_stupid_password
      login_user(user)
    end

    describe 'joinable_rooms' do
      it 'should receive Room.joinable' do
        expect(Room).to receive(:joinable).and_call_original
        post :service
      end
    end
  end
end
