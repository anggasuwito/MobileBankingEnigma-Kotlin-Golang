package com.app.livecode.mobilebanking.fragments

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.activityViewModels
import androidx.lifecycle.Observer
import androidx.recyclerview.widget.LinearLayoutManager
import com.app.livecode.R
import com.app.livecode.mobilebanking.recycleview.HistoryMobileBankingRecycleView
import com.app.livecode.mobilebanking.viewmodel.TransactionViewModel
import kotlinx.android.synthetic.main.fragment_history_mobile_banking.*

class HistoryMobileBankingFragment : Fragment() {
    val transactionViewModel by activityViewModels<TransactionViewModel>()
    lateinit var historyMobileBankingRecycleView:HistoryMobileBankingRecycleView
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_history_mobile_banking, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        historyTransactionRecycleView.layoutManager = LinearLayoutManager(this.context)
        transactionViewModel.listHistoryTransaction.observe(viewLifecycleOwner, Observer {
            historyMobileBankingRecycleView = HistoryMobileBankingRecycleView(it)
            historyTransactionRecycleView.adapter = historyMobileBankingRecycleView
        })
        transactionViewModel.getHistoryTransaction()
    }
}