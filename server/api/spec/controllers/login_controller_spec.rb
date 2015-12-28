require 'rails_helper'

RSpec.describe LoginController, type: :controller do
  context 'POST service' do
    let(:user) { create(:user, :with_stupid_password) }
    let(:params) { { name: user.name, password: 'secret' } }

    describe 'logged_in_user' do
      subject { @controller.logged_in_user }
      before { post :service, params }

      it { should be_a_kind_of(User) }
      it { should have_attributes(id: user.id, persisted?: true) }
    end

    describe 'joined_room' do
      subject { @controller.joined_room }

      context 'when the user has joined a room' do
        before do
          room = create(:room)
          room.join_user(user)
          post :service, params
        end

        it { should be_a_kind_of(Room) }
        it { should have_attributes(persisted?: true) }
      end

      context 'when the user has not joined a room' do
        before { post :service, params }
        it { should be_nil }
      end
    end
  end
end
