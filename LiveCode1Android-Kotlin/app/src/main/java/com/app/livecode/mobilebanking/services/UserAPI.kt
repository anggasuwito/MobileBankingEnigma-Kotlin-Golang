package com.app.livecode.mobilebanking.services

import com.app.livecode.mobilebanking.User
import retrofit2.Call
import retrofit2.http.GET

interface UserAPI {
    @GET("users")
    fun getUserInfo(): Call<List<User>>
}