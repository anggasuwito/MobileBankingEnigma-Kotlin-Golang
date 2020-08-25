package com.app.livecode.mobilebanking.services

import androidx.lifecycle.MutableLiveData
import com.app.livecode.mobilebanking.User
import com.google.gson.Gson
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response


class UserRepository(val userAPI: UserAPI) {
    var userInfo: MutableLiveData<List<User>> = MutableLiveData<List<User>>()

    fun getUserInfo() {
        userAPI.getUserInfo().enqueue(object : Callback<List<User>> {
            override fun onFailure(call: Call<List<User>>, t: Throwable) {
                println(t.localizedMessage)
                println(t.printStackTrace())
            }

            override fun onResponse(call: Call<List<User>>, response: Response<List<User>>) {
                val response = response.body()
                val gson = Gson()
                val stringResponse = gson.toJson(response)
                val userObject: List<User> =
                    gson.fromJson(stringResponse, Array<User>::class.java).toList()
                userInfo.value = userObject
            }
        })
    }
}