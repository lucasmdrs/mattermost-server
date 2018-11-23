// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// BotStore is an autogenerated mock type for the BotStore type
type BotStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: userId, includeDeleted
func (_m *BotStore) Get(userId string, includeDeleted bool) store.StoreChannel {
	ret := _m.Called(userId, includeDeleted)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(userId, includeDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields: page, perPage, includeDeleted
func (_m *BotStore) GetAll(page int, perPage int, includeDeleted bool) store.StoreChannel {
	ret := _m.Called(page, perPage, includeDeleted)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int, bool) store.StoreChannel); ok {
		r0 = rf(page, perPage, includeDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDelete provides a mock function with given fields: userId
func (_m *BotStore) PermanentDelete(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: bot
func (_m *BotStore) Save(bot *model.Bot) store.StoreChannel {
	ret := _m.Called(bot)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Bot) store.StoreChannel); ok {
		r0 = rf(bot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: bot
func (_m *BotStore) Update(bot *model.Bot) store.StoreChannel {
	ret := _m.Called(bot)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Bot) store.StoreChannel); ok {
		r0 = rf(bot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
