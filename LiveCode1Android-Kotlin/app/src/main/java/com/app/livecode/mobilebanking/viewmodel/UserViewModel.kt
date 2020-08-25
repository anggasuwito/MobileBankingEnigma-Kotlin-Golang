package com.app.livecode.mobilebanking.viewmodel

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import com.app.livecode.config.RetrofitBuilder
import com.app.livecode.mobilebanking.User
import com.app.livecode.mobilebanking.services.UserAPI
import com.app.livecode.mobilebanking.services.UserRepository

class UserViewModel : ViewModel() {
    val userRepository: UserRepository
    init {
        val userAPI = RetrofitBuilder.createRetrofit().create(UserAPI::class.java)
        userRepository =
            UserRepository(userAPI)
    }

    val userInfo: LiveData<List<User>> = userRepository.userInfo
    fun getUserInfo() {
        userRepository.getUserInfo()
    }
}