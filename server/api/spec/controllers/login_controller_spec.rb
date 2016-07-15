require 'rails_helper'

RSpec.describe LoginController, type: :controller do
  context 'POST service' do
    let(:user) { create(:user, :with_stupid_password) }
    let(:params) { { name: user.name, password: 'secret' } }

    it 'should call login' do
      expect(@controller).to receive(:login).and_call_original
      post :service, params: params
    end

    describe '#logged_in_user' do
      subject { assigns(:logged_in_user) }
      before do
        post :service, params: params
      end

      it { should be_a_kind_of(User) }
      its(:persisted?) { should be true }
    end
  end
end
