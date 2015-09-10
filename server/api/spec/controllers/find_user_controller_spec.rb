require 'rails_helper'

RSpec.describe FindUserController, type: :controller do
  context 'POST service' do
    before do
      user = create :user, :with_stupid_password
      login_user(user)

      create :user, name: 'Shimakaze'
      post :service, params
    end

    describe 'target_user' do
      context 'with an existing user name' do
        let(:params) { { name: 'Shimakaze' } }
        subject { @controller.target_user }
        it { should be_a_kind_of(User) }
        it { should have_attributes(persisted?: true, name: 'Shimakaze') }
      end

      context 'with a no-existing user name' do
        let(:params) { { name: 'Kaga' } }
        subject { assigns(:target_user) }
        it { should be_nil }
      end
    end
  end
end
