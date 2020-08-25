package com.app.livecode.mobilebanking.viewmodel

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import com.app.livecode.config.RetrofitBuilder
import com.app.livecode.mobilebanking.Transaction
import com.app.livecode.mobilebanking.User
import com.app.livecode.mobilebanking.services.TransactionAPI
import com.app.livecode.mobilebanking.services.TransactionRepository
import com.app.livecode.mobilebanking.services.UserAPI
import com.app.livecode.mobilebanking.services.UserRepository

class TransactionViewModel : ViewModel() {
    val transactionRepository: TransactionRepository

    init {
        val transactionAPI = RetrofitBuilder.createRetrofit().create(TransactionAPI::class.java)
        transactionRepository =
            TransactionRepository(transactionAPI)
    }
    val listHistoryTransaction : LiveData<List<Transaction>> = transactionRepository.historyTransaction
    fun getHistoryTransaction(){
        transactionRepository.getHistoryTransaction()
    }

    fun saveTransaction(transaction: Transaction){
        transactionRepository.createTransaction(transaction)
    }
}