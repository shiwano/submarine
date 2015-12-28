require 'rails_helper'

RSpec.describe CreateRoomController, type: :controller do
  context 'POST service' do
    before do
      user = create(:user, :with_stupid_password)
      login_user(user)
    end

    describe 'new_room' do
      it 'should return a new room' do
        post :service
        expect(@controller.new_room).to be_a_kind_of Room
      end
      it 'should receive current_user#create_room' do
        expect(@controller.current_user).to receive(:create_room).and_call_original
        post :service
      end
    end
  end
end
