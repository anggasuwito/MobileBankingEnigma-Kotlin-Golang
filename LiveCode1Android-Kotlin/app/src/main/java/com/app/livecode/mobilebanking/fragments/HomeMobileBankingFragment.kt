package com.app.livecode.mobilebanking.fragments

import android.content.Intent
import android.net.Uri
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.activityViewModels
import androidx.lifecycle.Observer
import androidx.navigation.NavController
import androidx.navigation.Navigation
import com.app.livecode.R
import com.app.livecode.mobilebanking.recycleview.HistoryMobileBankingRecycleView
import com.app.livecode.mobilebanking.viewmodel.TransactionViewModel
import com.app.livecode.mobilebanking.viewmodel.UserViewModel
import kotlinx.android.synthetic.main.fragment_home_mobile_banking.*
import kotlinx.android.synthetic.main.fragment_login_mobile_banking.*

class HomeMobileBankingFragment : Fragment(), View.OnClickListener {
    lateinit var navController: NavController
    val userViewModel by activityViewModels<UserViewModel>()
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_home_mobile_banking, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        println("ON VIEW CREATED")
        navController = Navigation.findNavController(view)
        userViewModel.getUserInfo()
        userViewModel.userInfo.observe(viewLifecycleOwner, Observer {
            userBalanceAmount.text = "$ " + it[0].balance
        })
        historyButton.setOnClickListener(this)
        callButton.setOnClickListener(this)
        transferButton.setOnClickListener(this)
    }

    override fun onResume() {
        super.onResume()
        println("ON RESUME")
        userViewModel.getUserInfo()
        userViewModel.userInfo.observe(viewLifecycleOwner, Observer {
            userBalanceAmount.text = "$ " + it[0].balance
        })
    }

    override fun onClick(v: View?) {
        when (v) {
            historyButton -> {
                navController.navigate(R.id.action_global_historyMobileBankingFragment)
            }
            transferButton -> {
                navController.navigate(R.id.actionHome_to_transferActivity)
            }
            callButton -> {
                val intent = Intent(Intent.ACTION_DIAL, Uri.parse("tel:085399921552"))
                startActivity(intent)
            }
        }
    }
}