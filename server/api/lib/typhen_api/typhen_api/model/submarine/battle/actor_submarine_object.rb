# This file was generated by typhen-api

module TyphenApi::Model::Submarine::Battle
  class ActorSubmarineObject
    include Virtus.model(:strict => true)

    attribute :equipment, TyphenApi::Model::Submarine::Battle::Equipment, :required => true
    attribute :is_using_pinger, Boolean, :required => true
  end
end
