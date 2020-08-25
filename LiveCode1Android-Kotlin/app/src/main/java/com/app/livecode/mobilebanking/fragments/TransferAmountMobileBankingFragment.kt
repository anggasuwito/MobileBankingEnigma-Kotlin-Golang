package com.app.livecode.mobilebanking.fragments

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Toast
import androidx.fragment.app.activityViewModels
import androidx.lifecycle.Observer
import androidx.navigation.NavController
import androidx.navigation.Navigation
import com.app.livecode.R
import com.app.livecode.mobilebanking.Transaction
import com.app.livecode.mobilebanking.User
import com.app.livecode.mobilebanking.viewmodel.TransactionViewModel
import com.app.livecode.mobilebanking.viewmodel.UserViewModel
import kotlinx.android.synthetic.main.fragment_home_mobile_banking.*
import kotlinx.android.synthetic.main.fragment_login_mobile_banking.*
import kotlinx.android.synthetic.main.fragment_transfer_amount_mobile_banking.*

class TransferAmountMobileBankingFragment : Fragment(), View.OnClickListener {
    lateinit var navController: NavController
    val transactionViewModel by activityViewModels<TransactionViewModel>()
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_transfer_amount_mobile_banking, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        navController = Navigation.findNavController(view)
        inputAmountButton.setOnClickListener(this)
    }

    override fun onClick(v: View?) {
        when (v) {
            inputAmountButton -> {
                if (amountRecipient.text.toString() == "" || descriptionRecipient.text.toString() == "") {
                    Toast.makeText(this.context, "Fill the form", Toast.LENGTH_SHORT).show()
                } else {
                    transactionViewModel.saveTransaction(Transaction(amount = amountRecipient.text.toString()))
                    navController.navigate(R.id.action_global_transferConfirmationMobileBankingFragment)
                }
            }
        }
    }
}