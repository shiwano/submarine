require 'rails_helper'

RSpec.describe SignUpController, type: :controller do
  context 'POST service' do
    let(:params) { { name: Faker::Name.first_name, password: Faker::Internet.password(6, 20) } }

    it 'should call login' do
      expect(@controller).to receive(:login).and_call_original
      post :service, params
    end

    describe 'new_user' do
      subject { assigns(:new_user) }
      before do
        post :service, params
      end

      it { should be_a_kind_of(User) }
      its(:persisted?) { should be true }
    end
  end
end
