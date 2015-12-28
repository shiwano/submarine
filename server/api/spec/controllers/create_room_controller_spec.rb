require 'rails_helper'

RSpec.describe CreateRoomController, type: :controller do
  context 'POST service' do
    before do
      user = create(:user, :with_stupid_password)
      login_user(user)
    end

    describe 'new_room' do
      subject { @controller.new_room }

      it { should be_a_kind_of(Room) }
      it 'should receive current_user#create_room' do
        expect(@controller.current_user).to receive(:create_room)
        subject
      end
    end
  end
end
